#include <GLFW/glfw3.h>

extern void onError(int, const char *);

// init initializes the error handling callback.
void init() {
	glfwSetErrorCallback(onError);
}
