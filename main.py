from ctypes import *

lib = cdll.LoadLibrary("lib/test.dll")

print(lib.hello("World"))

def logger(intVal, strVal):
    print("prefix:", intVal, strVal)
    return 0

CMPFUNC = CFUNCTYPE(c_int, c_int, c_wchar_p)
logger_func = CMPFUNC(logger)

print(lib.callback(logger_func))
