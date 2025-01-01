.data
.balign 8
main.4:
	.ascii "xdd: %f\n"
	.byte 0
/* end data */

.text
ahoj:
	pushq %rbp
	movq %rsp, %rbp
	movl %edi, %eax
	addl $9, %eax
	leave
	ret
.type ahoj, @function
.size ahoj, .-ahoj
/* end function ahoj */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	movss ".Lfp0"(%rip), %xmm0
	leaq main.4(%rip), %rdi
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
	.int 1085905306 /* 5.800000 */

.section .note.GNU-stack,"",@progbits
