section .text
global  _start

_start:
    mov eax, 2
    mov ebx, 3
    add eax, ebx

    mov ebx, eax            ; move result to ebx
    mov eax, 1              ; sys_exit
    int 0x80                ; calling sys_exit
