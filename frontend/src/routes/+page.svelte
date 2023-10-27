<script lang="ts">
	import OrgCard from "$lib/components/OrgCard.svelte";
	import { toastStyle } from "$lib/const.js";
	import { error } from "@sveltejs/kit";
	import { onMount } from "svelte";
	import toast, { Toaster } from "svelte-french-toast";

	export let data;


	onMount(() => {
		if (data.user == null) toast.error("Cannot load orgs informations", toastStyle)
	})

</script>

<Toaster/>
{#if data?.user}
	<div class="flex justify-center items-center w-full">
		<div class="grid grid-col-1 text-center">
			<h1 class="text-2xl text-center p-10">Bentornato {data.user?.display}</h1>
			{#if data.org}
				<div class="grid sm:grid-cols-[1fr_auto_1fr] grid-flow-row-dense grid-cols-1 gap-4	">
					{#each data.org as org}
						<OrgCard data={org}/>
					{/each}			
				</div>
			{:else}
				<p>No orgs found</p>
			{/if}
			<div class="flex h-full w-full justify-center items-center p-10">

				<a href="/logout" class="decoration-transparent"><button class="w-32 h-10 bg-gray-500 text-white">Logout</button></a>
			</div>

		</div>
	</div>
{/if}
