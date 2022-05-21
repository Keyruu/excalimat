<script lang="ts">
	import { page } from '$app/stores';

	function isStep(stepNumber: number) {
		let basePath = '/store';
		switch ($page.url.pathname) {
			case `${basePath}/user`:
				if (stepNumber === 1) return 'step-primary';
				break;
			case `${basePath}/login`:
				if (stepNumber <= 2) return 'step-primary';
				break;
			case `${basePath}/product`:
				if (stepNumber <= 3) return 'step-primary';
				break;
			case `${basePath}/buy`:
				if (stepNumber <= 4) return 'step-primary';
				break;
		}

		if ($page.url.pathname === `${basePath}user`) {
			if (stepNumber === 1) {
				return 'step-primary';
			}
		}
		if ($page.url.pathname === `${basePath}login`) {
		}
	}
</script>

<div class="navbar bg-base-100 flex">
	<a class="btn btn-ghost normal-case text-xl" href="/store/user">
		<img class="w-36" src="/excalimat_logo.png" alt="excalimat" />
	</a>
	<ul class="steps pt-2 flex-1 w-12">
		<li class={`step ${isStep(1)}`}>User auswählen</li>
		<li class={`step ${isStep(2)}`}>Einloggen</li>
		<li class={`step ${isStep(3)}`}>Produkt auswählen</li>
		<li class={`step ${isStep(4)}`}>Produkt kaufen</li>
	</ul>
</div>

<slot />
