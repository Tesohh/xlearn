<script lang="ts">
	import '../app.css';
	import { page } from '$app/stores';
	import type { User, Org } from '$lib/types';
	import IconHome from '$lib/components/home/IconHome.svelte';
	import OrgButton from '$lib/components/home/OrgButton.svelte';
	import { selectedOrg, selectedTab } from '$lib/writables';
	import { writable } from 'svelte/store';
	import JoinButton from '$lib/components/home/JoinButton.svelte';
	import { onMount } from 'svelte';

	export let data: { user: User; org: Org[] };

	let title = writable<string>('');

	$: selectedOrg.set($page.url.pathname.replace('/org/', ''));

	onMount(() => {
		let key = localStorage.getItem('tab');
		if (key != null) {
			selectedTab.set(key);
		}

		selectedTab.subscribe((u) => {
			if ($selectedTab != '') {
				localStorage.setItem('tab', u);
			}
		});
	});

	$: {
		const url = $page.url.pathname;

		if (url.startsWith('/org/')) {
			title.set(decodeURIComponent(`${url.replace('/org/', '')}`));
		} else if (url.startsWith('/login')) title.set('Login');
		else if (url.startsWith('/signup')) title.set('Signup');
		else if (url == '/') title.set('Home');
	}
</script>

<title>XLearn {$title != '' ? '- ' + $title : ''}</title>

{#if data.user}
	<div class="grid grid-cols-1 lg:grid-cols-[auto_440px] p-[18px] gap-[18px] h-screen">
		<div class="h-full border-primary border-2 rounded-md">
			<slot />
		</div>

		<div class="h-full hidden lg:block">
			<div class="flex flex-row h-[100px] border-primary border-2 rounded-md">
				<IconHome id="orgs" iconStr="icon-[tabler--smart-home]" />
				<IconHome id="shop" iconStr="icon-[tabler--shopping-bag]" />
				<IconHome id="test" iconStr="icon-[tabler--brand-docker]" />
				<IconHome id="user" iconStr="icon-[tabler--user-circle]" />
			</div>

			<div class="pt-[48px] grid grid-cols-1 gap-5">
				{#if $selectedTab == 'orgs'}
					{#if data.org}
						{#each data.org as org}
							<OrgButton data={org} />
						{/each}
						<JoinButton />
					{/if}
				{:else}
					<p>{$selectedTab}</p>
				{/if}
			</div>
		</div>
	</div>
{:else}
	<slot />
{/if}
