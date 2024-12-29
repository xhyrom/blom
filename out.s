.data
.balign 8
.0:
	.ascii "first: %f\n"
	.byte 0
/* end data */

.data
.balign 8
.1:
	.ascii "second: %f\n"
	.byte 0
/* end data */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	movsd ".Lfp1"(%rip), %xmm0
	leaq .0(%rip), %rdi
	movl $1, %eax
	callq printf
	movsd ".Lfp0"(%rip), %xmm0
	leaq .1(%rip), %rdi
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
	.int -1073741824
	.int 1077038284 /* 18.299999 */

.section .rodata
.p2align 3
.Lfp1:
	.int 0
	.int 1071644672 /* 0.500000 */

.section .note.GNU-stack,"",@progbits
