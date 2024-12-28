.data
.balign 8
main.0:
	.ascii "first: %f\n"
	.byte 0
/* end data */

.data
.balign 8
main.1:
	.ascii "second: %f\n"
	.byte 0
/* end data */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	movsd ".Lfp1"(%rip), %xmm0
	leaq main.0(%rip), %rdi
	movl $1, %eax
	callq printf
	movsd ".Lfp0"(%rip), %xmm0
	leaq main.1(%rip), %rdi
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
.p2align 3
.Lfp0:
	.int -858993459
	.int 1076022476 /* 9.400000 */

.section .rodata
.p2align 3
.Lfp1:
	.int 0
	.int 1071644672 /* 0.500000 */

.section .note.GNU-stack,"",@progbits
