<script lang="ts">
	import { toastStyle } from '$lib/const.js';
	import { onMount } from 'svelte';
	import toast from 'svelte-french-toast';
	import Icon from '@iconify/svelte';

	export let data;

	onMount(() => {
		if (data.error) toast.error(data.error, toastStyle);
	});
</script>

<div class="flex justify-center w-full text-center">
	<div class="grid grid-col-1 text-center items-center">
		<h1 class="font-semibold text-2xl p-10">{data.org?.name}</h1>

		<div class="grid grid-cols-1">
			<div class="flex flex-col items-center justify-center gap-20">
				{#if data.adventures}
					{#each data.adventures as adv}
						<h1>{JSON.stringify(adv)}</h1>
					{/each}
				{:else}
					<h1>Nothing to show here</h1>
				{/if}
			</div>
		</div>
	</div>
</div>
