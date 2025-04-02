<script lang="ts">
    import Counter from './Counter.svelte';
</script>

<main>
    <Counter />
</main>

<style>
    main {
        flex: 1;
        width: 100%;
        max-width: 800px;
        background: white;
        padding: 2rem;
        border-radius: 1rem;
        box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1);
        margin-bottom: 3rem;
    }
</style>