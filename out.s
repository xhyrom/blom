.data
.balign 8
main.0:
	.ascii "%d - %d\n"
	.byte 0
/* end data */

.data
.balign 8
main.1:
	.ascii "xdd: %f\n"
	.byte 0
/* end data */

.data
.balign 8
main.2:
	.ascii "gumi %d hm?\n"
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
	movl $9, %edi
	callq ahoj
	movl %eax, %edx
	movl $12, %esi
	leaq main.0(%rip), %rdi
	movl $0, %eax
	callq printf
	movsd ".Lfp0"(%rip), %xmm0
	leaq main.1(%rip), %rdi
	movl $1, %eax
	callq printf
	movl $9, %edi
	callq ahoj
	movl %eax, %esi
	leaq main.2(%rip), %rdi
	movl $0, %eax
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
	.int 858993459
	.int 1075262259 /* 5.800000 */

.section .note.GNU-stack,"",@progbits
