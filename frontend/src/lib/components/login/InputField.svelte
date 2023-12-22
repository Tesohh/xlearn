<script lang="ts">
	export let label: string;
	export let type: 'text' | 'password';
</script>

<label class="text-lg font-medium text-primary" for={label.toLowerCase().replaceAll(' ', '')}
	>{label}</label
>
<input
	required
	{type}
	name={label.toLowerCase().replaceAll(' ', '')}
	id={label.toLowerCase().replaceAll(' ', '')}
	class="border-2 text-md rounded-lg border-primary text-primary p-1"
/>
