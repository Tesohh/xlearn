<script lang="ts">
	import InputField from '$lib/components/login/InputField.svelte';
	import TabButton from '$lib/components/login/TabButton.svelte';
	import { toastStyle } from '$lib/const.js';
	import { tabState } from '$lib/writables';
	import { onMount } from 'svelte';
	import toast, { Toaster } from 'svelte-french-toast';
	import { writable } from 'svelte/store';

	onMount(() => {
		// @ts-ignore
		if (form?.from) tabState.set(form.from);
		if (form?.error) toast.error(form.error, toastStyle);
	});

	const zPressed = writable(0);

	function onKeydown(event: KeyboardEvent) {
		if (event.key == 'z') zPressed.update((n) => n + 1);
	}

	export let form;
</script>

<svelte:window on:keydown={onKeydown} />

<Toaster />

<div class="flex w-screen h-screen items-center justify-center">
	<div class="flex flex-col gap-0">
		<div class="w-96">
			<div
				class="grid grid-flow-col grid-rows-1 grid-cols-2 h-16 bg-primary rounded-t-lg rounded-b-none"
			>
				<TabButton title={'Login'} />
				<TabButton title={'Signup'} />
			</div>
			<div class="flex justify-center h-80 items-center border-primary border-4 rounded-b-lg">
				<!-- Login tab -->
				{#if $tabState == 'Login'}
					<form action="?/login" method="POST">
						<div class="flex flex-col items-center text-center justify-center">
							<InputField label={'Username'} type={'text'} />
							<InputField label={'Password'} type={'password'} />

							<div class="p-4">
								<button class="h-8 w-24 rounded-lg bg-primary text-secondary text-md" type="submit"
									>Login</button
								>
							</div>
						</div>
					</form>
					<!-- Sign up tab  -->
				{:else}
					<form action="?/signup" method="POST">
						<div class="flex flex-col items-center text-center justify-center">
							<InputField label={'Username'} type={'text'} />
							<InputField label={'Password'} type={'password'} />
							<InputField label={'Confirm Password'} type={'password'} />

							<div class="p-4">
								<button class="h-8 w-24 rounded-lg bg-primary text-secondary text-md" type="submit"
									>Signup</button
								>
							</div>
						</div>
					</form>
				{/if}
			</div>
		</div>
	</div>
</div>
