// X86_64 mock implementation - no LVGL dependencies
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <pthread.h>

// Mock types and constants
typedef void (*indev_handler_t)(int);
typedef void (*jetkvm_rpc_handler_t)(const char*, const char*);

static indev_handler_t indev_handler = NULL;
static jetkvm_rpc_handler_t rpc_handler = NULL;

// Mock UI variables
typedef struct {
    const char *name;
    void (*setter)(const char *);
    const char *(*getter)(void);
} ui_var_t;

static ui_var_t ui_vars[] = {
    // Add mock variables as needed
};
static int ui_vars_size = sizeof(ui_vars) / sizeof(ui_vars[0]);

// Mock implementations
void jetkvm_set_indev_handler(indev_handler_t handler) {
    indev_handler = handler;
    printf("[MOCK] Set input device handler\n");
}

const char *jetkvm_ui_event_code_to_name(int code) {
    static char buffer[32];
    snprintf(buffer, sizeof(buffer), "MOCK_EVENT_%d", code);
    return buffer;
}

void jetkvm_ui_set_var(const char *name, const char *value) {
    printf("[MOCK] UI set var: %s = %s\n", name, value);
}

const char *jetkvm_ui_get_var(const char *name) {
    printf("[MOCK] UI get var: %s\n", name);
    return "";
}

void jetkvm_ui_init(uint16_t rotation) {
    printf("[MOCK] UI init with rotation: %d\n", rotation);
}

void jetkvm_ui_tick() {
    // Mock UI tick - do nothing
}

void jetkvm_ui_set_rotation(uint16_t rotation) {
    printf("[MOCK] UI set rotation: %d\n", rotation);
}

const char *jetkvm_ui_get_current_screen() {
    return "main";
}

void jetkvm_ui_load_screen(const char *obj_name) {
    printf("[MOCK] UI load screen: %s\n", obj_name);
}

int jetkvm_ui_set_text(const char *obj_name, const char *text) {
    printf("[MOCK] UI set text: %s = %s\n", obj_name, text);
    return 0;
}

void jetkvm_ui_set_image(const char *obj_name, const char *image_name) {
    printf("[MOCK] UI set image: %s = %s\n", obj_name, image_name);
}

void jetkvm_ui_add_state(const char *obj_name, const char *state_name) {
    printf("[MOCK] UI add state: %s -> %s\n", obj_name, state_name);
}

void jetkvm_ui_clear_state(const char *obj_name, const char *state_name) {
    printf("[MOCK] UI clear state: %s -> %s\n", obj_name, state_name);
}

int jetkvm_ui_add_flag(const char *obj_name, const char *flag_name) {
    printf("[MOCK] UI add flag: %s -> %s\n", obj_name, flag_name);
    return 0;
}

int jetkvm_ui_clear_flag(const char *obj_name, const char *flag_name) {
    printf("[MOCK] UI clear flag: %s -> %s\n", obj_name, flag_name);
    return 0;
}

void jetkvm_ui_fade_in(const char *obj_name, uint32_t duration) {
    printf("[MOCK] UI fade in: %s (%u ms)\n", obj_name, duration);
}

void jetkvm_ui_fade_out(const char *obj_name, uint32_t duration) {
    printf("[MOCK] UI fade out: %s (%u ms)\n", obj_name, duration);
}

void jetkvm_ui_set_opacity(const char *obj_name, uint8_t opacity) {
    printf("[MOCK] UI set opacity: %s = %d\n", obj_name, opacity);
}

const char *jetkvm_ui_get_lvgl_version() {
    return "8.3.0-mock-x86_64";
}

// Mock RPC handler
void jetkvm_call_rpc_handler(const char *method, const char *params) {
    printf("[MOCK] RPC call: %s(%s)\n", method, params);
}

// Mock video functions
void jetkvm_video_init() {
    printf("[MOCK] Video init\n");
}

void jetkvm_video_shutdown() {
    printf("[MOCK] Video shutdown\n");
}

void jetkvm_video_start() {
    printf("[MOCK] Video start\n");
}

void jetkvm_video_stop() {
    printf("[MOCK] Video stop\n");
}

const char *jetkvm_video_log_status() {
    return "Mock video status: X86_64 simulation";
}

double jetkvm_video_get_stream_quality_factor() {
    return 1.0;
}

void jetkvm_video_set_stream_quality_factor(double factor) {
    printf("[MOCK] Video set quality factor: %f\n", factor);
}

const char *jetkvm_video_get_edid() {
    // Mock EDID for 1920x1080 display
    return "00ffffffffffff0010ac72404c384145"
           "2e120103802f1e78eaee95a3544c9926"
           "0f5054a54b00b300d100714fa9408180"
           "8140010101011d007251d01e206e2855"
           "00d9281100001e8c0ad08a20e02d1010"
           "3e9600138e2100001e023a8018713827"
           "40582c4500d9281100001e011d80d072"
           "1c1620102c2580d9281100009e000000";
}

void jetkvm_video_set_edid(const char *edid) {
    printf("[MOCK] Video set EDID: %s\n", edid);
}

// Mock crash function
void jetkvm_crash() {
    printf("[MOCK] Crash triggered\n");
    abort();
}

// Mock setup function
void jetkvm_setup_native_handlers() {
    printf("[MOCK] Setup native handlers for X86_64\n");
}