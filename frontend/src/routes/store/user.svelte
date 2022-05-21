<script lang="ts">
	import UserCard from '$lib/components/UserCard.svelte';
	import type { Account } from '$lib/types/Account';

	let accounts: Account[] = [];
	fetch('http://localhost:3000/excalimat/api/v1/account/')
		.then((res) => res.json())
		.then((json) => (accounts = json.data as Account[]));
</script>

<h1 class="h1">Login</h1>
<div class="flex justify-center align-middle flex-wrap">
	{#each accounts as account}
		<div class="p-2">
			<UserCard image={account.picture} name={account.name} balance={account.balance} />
		</div>
	{/each}
</div>
