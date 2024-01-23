<script lang="ts">
	import { toastStyle } from '$lib/const.js';
	import errorMessages from '$lib/errorMessages.js';
	import { onMount } from 'svelte';
	import toast, { Toaster } from 'svelte-french-toast';

	export let data;
	export let form;

	onMount(async () => {
		console.log(document.cookie);
		if (data.user == null) toast.error(errorMessages.orgsNotFound, toastStyle);
		if (form?.error) toast.error(errorMessages.errorWhileJoiningOrg, toastStyle);
	});

</script>

<Toaster />
{#if data?.user}
	<div class="flex justify-center items-center w-full">
		<div class="grid grid-col-1 text-center">
			<h1 class="text-2xl text-center p-10">Bentornato {data.user?.display}</h1>

			<a href="/logout" data-sveltekit-reload
				><button class="bg-gray-500 w-20 h-10 rounded-md">Logout</button></a
			>
		</div>
	</div>
{/if}
