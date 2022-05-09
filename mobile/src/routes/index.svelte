<script lang="ts">
</script>

<div class="card card-compact w-96 bg-base-100 shadow-xl">
	<figure><img src="https://api.lorem.space/image/shoes?w=400&h=225" alt="Shoes" /></figure>
	<div class="card-body">
		<h2 class="card-title">Shoes!</h2>
		<p>If a dog chews shoes whose shoes does he choose?</p>
		<div class="card-actions justify-end">
			<button class="btn btn-primary">Buy Now</button>
		</div>
	</div>
</div>
<h1 class="text-3xl font-bold underline">Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<style>
	:global(.switch) {
		position: relative;
		display: inline-flex;
		align-items: center;
		border-radius: 9999px;
		height: 1.5rem;
		width: 2.75rem;
	}

	:global(.switch-enabled) {
		/* Blue */
		background-color: rgb(37 99 235);
	}

	:global(.switch-disabled) {
		/* Gray */
		background-color: rgb(229 231 235);
	}
</style>
