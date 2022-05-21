<script lang="ts">
	import UserCard from '$lib/components/UserCard.svelte';
</script>

<h1 class="h1">Login</h1>
<div class="flex justify-center align-middle flex-wrap">
	{#each Array(9) as _, i}
		<div class="p-2">
			<UserCard image={'https://api.lorem.space/image/face?hash=' + i} name="Lucas Rott" />
		</div>
	{/each}
</div>
