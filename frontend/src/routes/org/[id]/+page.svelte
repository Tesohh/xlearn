<script lang="ts">
	import Adventure from '$lib/components/org/Adventure.svelte';
	import { toastStyle } from '$lib/const.js';
	import errorMessages from '$lib/errorMessages.js';
	import { onMount } from 'svelte';
	import toast from 'svelte-french-toast';

	export let data;

	onMount(() => {
		if (data.error) toast.error(data.error, toastStyle);
	});
</script>

<div class="flex justify-center w-full text-center">
	<div class="grid grid-col-1 text-center items-center">
		<div class="grid grid-cols-1">
			<div class="flex flex-col items-center justify-center gap-20 py-5">
				{#if data.adventures}
					{#each data.adventures as adv}
						<Adventure data={adv} />
					{/each}
				{:else}
					<h1>{errorMessages.anyAdventureLoaded}</h1>
				{/if}
			</div>
		</div>
	</div>
</div>
