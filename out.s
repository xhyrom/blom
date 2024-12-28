.data
.balign 8
guhu.0:
	.ascii "lol"
	.byte 0
/* end data */

.data
.balign 8
main.1:
	.ascii "ccc: %f\n"
	.byte 0
/* end data */

.data
.balign 8
main.2:
	.ascii "lol: %d\n"
	.byte 0
/* end data */

.data
.balign 8
main.3:
	.ascii "lul: %d\n"
	.byte 0
/* end data */

.data
.balign 8
main.4:
	.ascii "lal: %s\n"
	.byte 0
/* end data */

.text
abc:
	pushq %rbp
	movq %rsp, %rbp
	movl $12, %eax
	leave
	ret
.type abc, @function
.size abc, .-abc
/* end function abc */

.text
guhu:
	pushq %rbp
	movq %rsp, %rbp
	leaq guhu.0(%rip), %rax
	leave
	ret
.type guhu, @function
.size guhu, .-guhu
/* end function guhu */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	pushq %rbx
	pushq %r12
	callq abc
	movzbl %al, %r12d
	callq guhu
	movq %rax, %rbx
	movsd ".Lfp0"(%rip), %xmm0
	leaq main.1(%rip), %rdi
	movl $1, %eax
	callq printf
	movl %r12d, %esi
	movl %esi, %r12d
	movl $600, %esi
	leaq main.2(%rip), %rdi
	movl $0, %eax
	callq printf
	movl %r12d, %esi
	leaq main.3(%rip), %rdi
	movl $0, %eax
	callq printf
	movq %rbx, %rsi
	leaq main.4(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $0, %eax
	popq %r12
	popq %rbx
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

/* floating point constants */
.section .rodata
.p2align 3
.Lfp0:
	.int 0
	.int 1078689792 /* 55.000000 */

.section .note.GNU-stack,"",@progbits
