.data
.balign 8
ahoj.0:
	.ascii "Guten Tag!"
	.byte 0
/* end data */

.data
.balign 8
main.1:
	.ascii "%d - %s\n"
	.byte 0
/* end data */

.data
.balign 8
main.2:
	.ascii "gumi %f hm?\n"
	.byte 0
/* end data */

.text
ahoj:
	pushq %rbp
	movq %rsp, %rbp
	leaq ahoj.0(%rip), %rax
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
	callq ahoj
	movq %rax, %rdx
	movl $12, %esi
	leaq main.1(%rip), %rdi
	movl $0, %eax
	callq printf
	movsd ".Lfp0"(%rip), %xmm0
	leaq main.2(%rip), %rdi
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
	.int -1717986918
	.int 1075157401 /* 5.400000 */

.section .note.GNU-stack,"",@progbits
