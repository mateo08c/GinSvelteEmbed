<script lang="ts">
    import {onMount} from 'svelte';
    import {marked} from 'marked';

    let readmeContent: string | Promise<string> = '';
    let loading = true;
    let error: string | null = null;

    async function fetchReadme() {
        try {
            const response = await fetch('https://raw.githubusercontent.com/mateo08c/GinSvelteEmbed/main/README.md');
            if (!response.ok) throw new Error('Failed to fetch README');
            const text = await response.text();
            readmeContent = marked(text);
            error = null;
        } catch (e) {
            error = e instanceof Error ? e.message : 'Failed to load README';
        } finally {
            loading = false;
        }
    }

    onMount(fetchReadme);
</script>

<section class="readme">
    <div class="readme-content">
        {#if loading}
            <div class="loading">Loading README...</div>
        {:else if error}
            <div class="error">
                Failed to load README: {error}
            </div>
        {:else}
            <article>
                {@html readmeContent}
            </article>
        {/if}
    </div>
</section>

<style>
    .readme {
        width: 100%;
        max-width: 800px;
        background: white;
        border-radius: 1rem;
        box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1);
        padding: 0 2rem 2rem;
    }

    .readme-content {
        color: var(--vp-c-text-1);
    }

    .readme-content :global(h1) {
        font-size: 2rem;
        margin: 2rem 0 1rem;
    }

    .readme-content :global(h2) {
        font-size: 1.5rem;
        margin: 1.5rem 0 1rem;
    }

    .readme-content :global(h3) {
        font-size: 1.2rem;
        margin: 1.2rem 0 0.8rem;
    }

    .readme-content :global(p) {
        margin: 1rem 0;
        line-height: 1.6;
    }

    .readme-content :global(ul) {
        list-style-type: disc;
        padding-left: 1.5rem;
        margin: 1rem 0;
    }

    .readme-content :global(li) {
        margin: 0.5rem 0;
    }

    .readme-content :global(pre) {
        background: #f1f5f9;
        padding: 1rem;
        border-radius: 0.5rem;
        overflow-x: auto;
        margin: 1rem 0;
    }

    .readme-content :global(code) {
        font-family: 'Fira Code', monospace;
        font-size: 0.9em;
    }

    .loading {
        text-align: center;
        color: var(--vp-c-text-2);
        padding: 2rem;
    }

    .error {
        color: var(--vp-c-danger);
        background-color: var(--vp-c-danger-soft);
        padding: 1rem;
        border-radius: 0.5rem;
        margin: 1rem 0;
    }
</style>