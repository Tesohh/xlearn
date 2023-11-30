<script lang='ts'>
	import type { OrgType } from "$lib/types";

    export let data: OrgType;
</script>

<div>
    <a href="/org/@{data.name}" class="decoration-transparent">
        <div class="flex justify-center items-center text-center bg-gray-400 h-28">
            <div class="grid grid-cols-1 gap-2">
                <h1 class="w-48">{data.name}</h1>
                {#if data.adventures}
                    <p>{data.adventures?.length} adventures</p>
                {/if}
                
            </div>
    
        </div>
    </a>

    <!-- Leave org button -->
    <div class="flex bottom-2 bg-red-400 h-8 justify-center text-center">
        <form action="?/leaveorg" method="post">
            <!-- In order to send to the server the org name -->
            <input type="text" name="code" placeholder={data.name} class="hidden">
            <button type="submit">Leave org</button>
        </form>
    </div>

</div>