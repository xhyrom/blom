.data
.balign 8
s.2a:
	.ascii "abc"
	.byte 0
/* end data */

.data
.balign 8
s.2:
	.ascii "lol: %d"
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
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	leaq s.2a(%rip), %rdi
	callq puts
	movl $9, %esi
	leaq s.2(%rip), %rdi
	movl $0, %eax
	callq printf
	movl $0, %eax
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
