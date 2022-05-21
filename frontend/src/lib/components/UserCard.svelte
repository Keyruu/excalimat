<script lang="ts">
	export let image: string;
	export let name: string;
</script>

<div class="card w-44 bg-secondary shadow-xl h-56 btn normal-case">
	<figure class="px-2 pt-6">
		<div class="avatar">
			<div class="w-28 mask mask-squircle">
				<img src={image} />
			</div>
		</div>
	</figure>
	<div class="card-body items-center text-center p-4">
		<h2 class="card-title p">{name}</h2>
		<p>-20,00$</p>
	</div>
</div>
