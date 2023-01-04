<script>
	import { createEventDispatcher, onDestroy, onMount } from 'svelte';

	const dispatch = createEventDispatcher();
	const close = () => dispatch('close');

	let modal;

	onMount(() => {
		disableScrolling();
		modal.focus();
	});

	onDestroy(() => {
		enableScrolling();
	});

	const handle_keydown = e => {
		if (e.key === 'Escape') {
			close();
			return;
		}

		if (e.key === 'Tab') {
			// trap focus
			const nodes = modal.querySelectorAll('*');
			const tabbable = Array.from(nodes).filter(n => n.tabIndex >= 0);

			let index = tabbable.indexOf(document.activeElement);
			if (index === -1 && e.shiftKey) index = 0;

			index += tabbable.length + (e.shiftKey ? -1 : 1);
			index %= tabbable.length;

			tabbable[index].focus();
			e.preventDefault();
		}
	};

	function disableScrolling(){
		var x=window.scrollX;
		var y=window.scrollY;
		window.onscroll=function(){window.scrollTo(x, y);};
	}

	function enableScrolling(){
		window.onscroll=function(){};
	}
</script>

<svelte:window on:keydown={handle_keydown}/>

<div class="modal-background" on:click={close} on:keydown={handle_keydown}></div>

<div class="modal" role="dialog" aria-modal="true" bind:this={modal}>
	<!-- <span class="closeButtonPlacement"><button class="close" on:click={close}>X</button></span> -->
	<slot name="header"></slot>
	<hr>
	<slot name="body"></slot>
	<hr>
	<slot name="footer"></slot>
</div>

<style lang="scss">
	.modal-background {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0,0,0,0.3);
	}

	.modal {
		position: fixed;
		left: 50vw;
		top: 50vh;
		width: calc(100vw - 4em);
		max-width: 32em;
		max-height: calc(100vh - 4em);
		overflow: auto;
		transform: translate(-50%,-50%);
		padding: 1em;
		border-radius: 1em;
        background-color: #383a3c
	}
</style>