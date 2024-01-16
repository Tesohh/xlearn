<script lang="ts">
	export let iconStr: string;
	export let url: string;
</script>

<a href={url}>
	<div class="h-full w-[110px] flex items-center justify-center">
		<svg class="{iconStr} h-[50px] w-[50px] bg-primary text-sm" />
	</div>
</a>
