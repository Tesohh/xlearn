<script lang='ts'>
	import { toastStyle } from '$lib/const.js';
	import { onMount } from 'svelte';
	import toast, { Toaster } from 'svelte-french-toast';

    export let form

    onMount(() => {
        if (form?.error) toast.error("Error while joining. Retry", toastStyle)
    })

</script>

<Toaster/>

<div class="flex justify-center items-center w-screen h-screen">

    <div class="grid grid-cols-1 grid-rows-2 text-center justify-center">

        <h1>Join org</h1>
        <form action="?/joinorg" method="POST" on:submit={() => { toast.loading("Joining in progress", toastStyle) }}>
            <button type="submit" class="p-2 bg-gray-500 rounded-md w-24">Join</button>
        </form>

    </div>

</div>