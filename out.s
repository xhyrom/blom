.data
.balign 8
main.6:
	.ascii "%f\n"
	.byte 0
/* end data */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	movss ".Lfp0"(%rip), %xmm0
	leaq main.6(%rip), %rdi
	movl $1, %eax
	callq printf
	movl $0, %eax
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

/* floating point constants */
.section .rodata
.p2align 2
.Lfp0:
	.int 1092091904 /* 9.500000 */

.section .note.GNU-stack,"",@progbits
