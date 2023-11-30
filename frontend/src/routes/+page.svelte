<script lang="ts">
	import { invalidate } from "$app/navigation";
	import OrgCard from "$lib/components/OrgCard.svelte";
	import { toastStyle } from "$lib/const.js";
	import errorMessages from "$lib/errorMessages.js";
	import { onMount } from "svelte";
	import toast, { Toaster } from "svelte-french-toast";

	export let data;
	export let form;

	onMount(async () => {
		if (data.user == null) toast.error(errorMessages.orgsNotFound, toastStyle)
		if (form?.error) toast.error(form.error, toastStyle)
	})

</script>

<Toaster/>
{#if data?.user}
	<div class="flex justify-center items-center w-full">

		<div class="grid grid-col-1 text-center">

			<h1 class="text-2xl text-center p-10">Bentornato {data.user?.display}</h1>

			{#if data.org}

				<div class="grid sm:grid-cols-[1fr_auto_1fr] grid-flow-row-dense grid-cols-1 gap-4">
					{#each data.org as org}
						<OrgCard data={org}/>
					{/each}
				</div>

			{:else}
				<p>No orgs found</p>
			{/if}
			<!-- Joining org form -->
			<div class="flex h-full w-full justify-center items-center p-10">
				
				<form method="post" action="?/joinorg">
					<input type="text" name="code" class="border-black border-2 h-10 rounded-md">
					<button type="submit" class="bg-gray-500 w-20 h-10 rounded-md">Join</button>
				</form>
				
			</div>
		</div>
	</div>
{/if}