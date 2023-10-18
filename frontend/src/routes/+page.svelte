<script lang="ts">
	export let data;
</script>

<style>
	td {
		border-width: 1px;
		border-color: black;
		border-collapse: collapse;
		padding: 5px;
	}
</style>

<div class="flex justify-center items-center w-screen p-5">
	<div class="grid grid-rows-2 grid-col-1">
		<h1 class="text-2xl text-center">
			<p>Informazioni utente</p>
			<table class="items-center border-collapse border border-black p-10">
				<tr>
					<td>Username</td>
					<td>Display Name</td>
					<td>Coins</td>
					<td>Ruolo</td>
				</tr>
				<tr>
					<td>{data.username}</td>
					<td>{data.display}</td>
					<td>{data.coins}</td>
					<td>{data.role == 1 ? "Studente" : data.role == 2 ? "Insegnante" : data.role == 3 ? "Admin" : "Not found"}</td>
				</tr>
			</table>
		</h1>
	
		<div class="flex h-full w-full justify-center items-center ">

			<a href="/logout" class="decoration-transparent"><button class="w-32 h-10 bg-gray-500 text-white" >Logout</button></a>
		</div>

	</div>


</div>
