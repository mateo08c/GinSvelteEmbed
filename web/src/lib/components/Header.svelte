<script lang="ts">
    import {Github} from 'lucide-svelte';
</script>

<header>
    <div class="header-content">
        <h1>Go + Svelte Counter</h1>
        <p class="description">
            A seamless integration between Svelte and Golang.</p>
        <div class="actions">
            <a
                    href="https://github.com/mateo08c/GinSvelteEmbed"
                    target="_blank"
                    rel="noopener"
                    class="github-link"
                    aria-label="View source code on GitHub"
            >
                <Github size={20}/>
                <span>View on GitHub</span>
            </a>
        </div>
    </div>
    <div class="header-decoration"></div>
</header>

<style>
    header {
        text-align: center;
        margin-bottom: 3rem;
    }

    h1 {
        font-size: 2.5rem;
        color: var(--vp-c-text-1);
        margin-bottom: 1rem;
    }

    .description {
        color: var(--vp-c-text-2);
        max-width: 600px;
        margin: 0 auto;
        line-height: 1.6;
    }

    .description {
        margin-bottom: 1.5rem;
    }

    .github-link {
        display: inline-flex;
        align-items: center;
        gap: 0.5rem;
        padding: 0.5rem 1rem;
        background-color: var(--vp-c-brand);
        color: white;
        border-radius: 0.5rem;
        text-decoration: none;
        transition: background-color 0.2s;
    }

    .github-link:hover {
        background-color: var(--vp-c-brand-dark);
        text-decoration: none;
    }
</style>