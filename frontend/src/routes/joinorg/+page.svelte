<script lang='ts'>
	import { toastStyle } from '$lib/const.js';
	import { error } from '@sveltejs/kit';
	import { onMount } from 'svelte';
	import toast, { Toaster } from 'svelte-french-toast';

    export let data

    onMount(() => {
        if (data?.error) toast.error(data.error, toastStyle)
    })

</script>

<Toaster/>

{#if data.error == null}
<div class="flex justify-center items-center w-screen h-screen">
    <div class="grid grid-cols-1 grid-rows-2 text-center justify-center">
        <h1>Welcome to {data.org}</h1>
        <a href="/"><button type="submit" class="p-2 bg-gray-500 rounded-md w-36">Start learning</button></a>
    </div>
</div>
{/if}