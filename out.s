.data
.balign 8
.0:
	.ascii "e: %f\n"
	.byte 0
/* end data */

.text
fact:
	pushq %rbp
	movq %rsp, %rbp
	subq $8, %rsp
	pushq %rbx
	cmpq $1, %rdi
	jle .Lbb2
	movq %rdi, %rbx
	subq $1, %rdi
	callq fact
	movq %rbx, %rdi
	imulq %rdi, %rax
	jmp .Lbb3
.Lbb2:
	movl $1, %eax
.Lbb3:
	popq %rbx
	leave
	ret
.type fact, @function
.size fact, .-fact
/* end function fact */

.text
eulersnumber:
	pushq %rbp
	movq %rsp, %rbp
	subq $16, %rsp
	pushq %rbx
	pushq %r12
	movsd ".Lfp0"(%rip), %xmm15
	movsd %xmm15, -16(%rbp)
	movl $0, %ebx
.Lbb6:
	cmpq $50, %rbx
	jge .Lbb9
	subq $16, %rsp
	movq %rsp, %r12
	movq %rbx, %rdi
	callq fact
	movq %rax, %rcx
	movsd -16(%rbp), %xmm0
	movq %rcx, (%r12)
	subq $16, %rsp
	movq %rsp, %rax
	cvtsi2sd %rcx, %xmm1
	movsd %xmm1, %xmm15
	movsd ".Lfp1"(%rip), %xmm1
	divsd %xmm15, %xmm1
	addsd %xmm1, %xmm0
	movsd %xmm0, (%rax)
	subq $16, %rsp
	movq %rsp, %rax
	addq $1, %rbx
	movq %rbx, (%rax)
	movsd %xmm0, -16(%rbp)
	jmp .Lbb6
.Lbb9:
	movsd -16(%rbp), %xmm0
	movq %rbp, %rsp
	subq $32, %rsp
	popq %r12
	popq %rbx
	leave
	ret
.type eulersnumber, @function
.size eulersnumber, .-eulersnumber
/* end function eulersnumber */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	callq eulersnumber
	leaq .0(%rip), %rdi
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
	.int 0
	.int 0 /* 0.000000 */

.section .rodata
.p2align 3
.Lfp1:
	.int 0
	.int 1072693248 /* 1.000000 */

.section .note.GNU-stack,"",@progbits
