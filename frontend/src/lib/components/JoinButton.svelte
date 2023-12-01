<script lang='ts'>
	import { joinPopupTrigger } from '$lib/writables';
    import { IconPlus } from '@tabler/icons-svelte';

    function setTrigger() {
        if ($joinPopupTrigger == true) joinPopupTrigger.set(false)
        else joinPopupTrigger.set(true)
    }

</script>

<div class="flex flex-row items-center">
    <button on:click={setTrigger}>
        <div class="w-12 h-12 bg-blue-800 rounded-md hover:bg-blue-400 hover:rounded-sm">
    
            <div class="flex justify-center items-center w-full h-full group">
                <IconPlus color="white" />
    
                {#if $joinPopupTrigger == false}
                    <span class="absolute w-auto p-2 m-2 min-w-max left-16 rounded-md shadow-md text-white bg-blue-700 font-bold scale-0 group-hover:scale-100">Join a new org</span>
                {/if}
    
            </div>
            
        </div>
    </button>

    {#if $joinPopupTrigger}
        
        <div class="absolute left-20 bg-gray-500 rounded-md">
            <div class="flex h-full w-full justify-center items-center p-10">
                
                <form method="post" action="?/joinorg">
                    <input type="text" name="code" class="border-black border-2 h-10 rounded-md" placeholder="Join code">
                    <button type="submit" data-sveltekit-reload class="bg-gray-700 w-20 h-10 rounded-md">Join</button>
                </form>
                
                
            </div>

        </div>
    {/if}

</div>