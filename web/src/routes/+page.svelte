<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import CounterSection from '$lib/components/CounterSection.svelte';
	import ReadmeSection from '$lib/components/ReadmeSection.svelte';
	import Footer from '$lib/components/Footer.svelte';
</script>

<div class="container">
	<Header />
	<CounterSection />
	<ReadmeSection />
	<Footer />
</div>

<style>
	:global(:root) {
		--vp-c-brand: #646cff;
		--vp-c-brand-dark: #535bf2;
		--vp-c-text-1: #213547;
		--vp-c-text-2: #666;
		--vp-c-danger: #dc2626;
		--vp-c-danger-soft: #fee2e2;
	}

	:global(body) {
		margin: 0;
		background-color: #f9fafb;
		color: var(--vp-c-text-1);
	}

	.container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 2rem;
		min-height: 100vh;
		display: flex;
		flex-direction: column;
		align-items: center;
	}
</style>