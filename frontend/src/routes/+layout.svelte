<script lang="ts">
	import Button from '$lib/components/Button.svelte';
		import type { UserType, OrgType } from '$lib/types';
	import '../app.css';

	export let data : { user: UserType, org: OrgType[] };

</script>

{#if data.user}
	<div class="grid md:grid-cols-[64px_auto_64px] grid-rows-1 gap-0">

		<!-- Left side bar -->
		<div class="bg-blue-300 w-16 h-screen hidden md:block">
			<div class="flex flex-col items-center justify-center py-10 gap-10">
				{#if data.org}

					{#each data.org as org}
						<Button data={org}/>
					{/each}

				{/if}
				<Button data={null}/>

			</div>
		</div>
		
		<!-- Main content  -->
		<div class="w-full h-full">
			<slot />
		</div>
		
		<!-- Righ side bar -->
		<div class="bg-blue-300 w-16 absolute right-0 h-screen hidden md:block">
			<div class="flex flex-col items-center justify-center py-10 gap-10">
				<!-- <Button data={"Test1"} /> -->
				<!-- <Button data={"Test2"} /> -->
				<!-- <Button data={"Test3"} /> -->

			</div>
		</div>

	</div>
{:else}
	<slot/>
{/if}
