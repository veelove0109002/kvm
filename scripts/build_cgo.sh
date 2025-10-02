#!/bin/bash
set -e

SCRIPT_PATH=$(realpath "$(dirname $(realpath "${BASH_SOURCE[0]}"))")
source ${SCRIPT_PATH}/build_utils.sh

CGO_PATH=$(realpath "${SCRIPT_PATH}/../internal/native/cgo")
BUILD_DIR=${CGO_PATH}/build

# Determine toolchain based on architecture
if [ "${TARGET_ARCH}" = "x86_64" ]; then
    CMAKE_TOOLCHAIN_FILE=""
    CMAKE_SYSTEM_PROCESSOR="x86_64"
    CMAKE_CROSSCOMPILING=0
else
    CMAKE_TOOLCHAIN_FILE=/opt/jetkvm-native-buildkit/rv1106-jetkvm-v2.cmake
    CMAKE_SYSTEM_PROCESSOR="armv7l"
    CMAKE_CROSSCOMPILING=1
fi

CLEAN_ALL=${CLEAN_ALL:-0}

if [ "$CLEAN_ALL" -eq 1 ]; then
    rm -rf "${BUILD_DIR}"
fi

TMP_DIR=$(mktemp -d)
pushd "${CGO_PATH}" > /dev/null

msg_info "▶ Generating UI index"
./ui_index.gen.sh

msg_info "▶ Building native library"
if [ "${TARGET_ARCH}" = "x86_64" ]; then
    # Use X86_64 specific CMakeLists.txt
    cp "${CGO_PATH}/CMakeLists_x86_64.txt" "${CGO_PATH}/CMakeLists.txt.bak"
    mv "${CGO_PATH}/CMakeLists.txt" "${CGO_PATH}/CMakeLists_original.txt" 2>/dev/null || true
    cp "${CGO_PATH}/CMakeLists_x86_64.txt" "${CGO_PATH}/CMakeLists.txt"
    
    # X86_64 native build
    VERBOSE=1 cmake -B "${BUILD_DIR}" \
        -DCMAKE_SYSTEM_PROCESSOR=${CMAKE_SYSTEM_PROCESSOR} \
        -DCMAKE_SYSTEM_NAME=Linux \
        -DCMAKE_CROSSCOMPILING=${CMAKE_CROSSCOMPILING} \
        -DLV_BUILD_USE_KCONFIG=ON \
        -DLV_BUILD_DEFCONFIG_PATH=${CGO_PATH}/lvgl_defconfig \
        -DCONFIG_LV_BUILD_EXAMPLES=OFF \
        -DCONFIG_LV_BUILD_DEMOS=OFF \
        -DSKIP_GLIBC_NAMES=ON \
        -DCMAKE_BUILD_TYPE=Release \
        -DCMAKE_INSTALL_PREFIX="${TMP_DIR}" \
        -DTARGET_ARCH=x86_64
else
    # ARM cross-compilation build
    VERBOSE=1 cmake -B "${BUILD_DIR}" \
        -DCMAKE_SYSTEM_PROCESSOR=${CMAKE_SYSTEM_PROCESSOR} \
        -DCMAKE_SYSTEM_NAME=Linux \
        -DCMAKE_CROSSCOMPILING=${CMAKE_CROSSCOMPILING} \
        -DCMAKE_TOOLCHAIN_FILE=$CMAKE_TOOLCHAIN_FILE \
        -DLV_BUILD_USE_KCONFIG=ON \
        -DLV_BUILD_DEFCONFIG_PATH=${CGO_PATH}/lvgl_defconfig \
        -DCONFIG_LV_BUILD_EXAMPLES=OFF \
        -DCONFIG_LV_BUILD_DEMOS=OFF \
        -DSKIP_GLIBC_NAMES=ON \
        -DCMAKE_BUILD_TYPE=Release \
        -DCMAKE_INSTALL_PREFIX="${TMP_DIR}"
fi

msg_info "▶ Copying built library and header files"
cmake --build "${BUILD_DIR}" --target install
cp -r "${TMP_DIR}/include" "${CGO_PATH}"
cp -r "${TMP_DIR}/lib" "${CGO_PATH}"
rm -rf "${TMP_DIR}"

# Restore original CMakeLists.txt if we modified it
if [ "${TARGET_ARCH}" = "x86_64" ] && [ -f "${CGO_PATH}/CMakeLists_original.txt" ]; then
    mv "${CGO_PATH}/CMakeLists_original.txt" "${CGO_PATH}/CMakeLists.txt"
fi

popd > /dev/null
