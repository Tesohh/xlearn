<script lang="ts">
	import { page } from '$app/stores';
	import JoinButton from '$lib/components/JoinButton.svelte';
	import OrgButton from '$lib/components/OrgButton.svelte';
	import type { UserType, OrgType } from '$lib/types';
	import { onMount } from 'svelte';
	import '../app.css';
	import HomeButton from '$lib/components/HomeButton.svelte';


	export let data : { user: UserType, org: OrgType[] };

	let title = ""

	onMount(() => {
		dynamicTitle()
	})

	function dynamicTitle() {

		const url = $page.url.pathname

		if (url.startsWith("/org/")) { title = decodeURIComponent(`${url.replace("/org/", "")}`) } 
		else if (url.startsWith("/login") ) title = "Login"
		else if (url.startsWith("/signup") ) title = "Signup"
		else if (url == "/") title = "Home"

	}

</script>

<title>XLearn {title != "" ? "- " + title : ""}</title>

{#if data.user}
		<div class="grid md:grid-cols-[64px_auto_64px] grid-rows-1 gap-0">
	
			<!-- Left side bar -->
			<div class="bg-blue-300 w-16 h-screen hidden md:block py-10 fixed bottom-0 top-0 z-10">
				<div class="grid grid-rows-1 grid-flow-row gap-10 justify-center">
					<HomeButton />
					{#if data.org}
	
						{#each data.org as org}
							<OrgButton data={org}/>
						{/each}
	
					{/if}
					<JoinButton />
				</div>
			</div>
			
			<!-- Main content  -->
			<div class="w-full h-full absolute z-0">
				<slot/>
			</div>
			
			<!-- Righ side bar -->
			<div class="bg-blue-300 w-16 right-0 h-screen hidden md:block fixed z-10">
				<div class="flex flex-col items-center justify-center py-10 gap-10">
	
				</div>
			</div>
	
		</div>
	{:else}

	<slot/>
{/if}
