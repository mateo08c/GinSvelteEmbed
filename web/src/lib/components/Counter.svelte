<script lang="ts">
    import {onMount} from 'svelte';
    import {Minus, Plus} from 'lucide-svelte';
    import {fade} from 'svelte/transition';

    let count = $state(0)
    let error: string | null = $state(null);
    let loading = $state(false);

    async function fetchCount() {
        try {
            loading = true;
            const response = await fetch('/api/count');
            if (!response.ok) throw new Error('Failed to fetch count');
            const data = await response.json();
            count = data.count;
            error = null;
        } catch (e) {
            error = e instanceof Error ? e.message : 'An error occurred';
        } finally {
            loading = false;
        }
    }

    async function updateCount(action: 'increment' | 'decrement') {
        try {
            loading = true;
            const response = await fetch(`/api/${action}`, {method: 'POST'});
            if (!response.ok) throw new Error(`Failed to ${action}`);
            const data = await response.json();
            count = data.count;
            error = null;
        } catch (e) {
            error = e instanceof Error ? e.message : 'An error occurred';
        } finally {
            loading = false;
        }
    }

    onMount(fetchCount);
</script>

<div class="counter-container">
    {#if error}
        <div class="error" transition:fade>
            {error}
        </div>
    {/if}

    <div class="counter-display" class:loading>
         <span class="count">
             {count}
         </span>
    </div>

    <div class="button-group">
        <button
                onclick={() => updateCount('decrement')}
                disabled={loading}
                class="counter-button decrement"
                aria-label="Décrémenter"
        >
            <Minus size={20}/>
        </button>
        <button
                onclick={() => updateCount('increment')}
                disabled={loading}
                class="counter-button increment"
                aria-label="Incrémenter"
        >
            <Plus size={20}/>
        </button>
    </div>
</div>

<style>
    .counter-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 1.5rem;
    }

    .counter-display {
        font-size: 4rem;
        font-weight: bold;
        color: var(--vp-c-text-1);
        min-height: 6rem;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .counter-display.loading {
        opacity: 0.5;
    }

    .button-group {
        display: flex;
        gap: 1rem;
    }

    .counter-button {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 3rem;
        height: 3rem;
        border-radius: 0.5rem;
        background-color: var(--vp-c-brand);
        color: white;
        border: none;
        cursor: pointer;
        transition: background-color 0.2s;
    }

    .counter-button:hover {
        background-color: var(--vp-c-brand-dark);
    }

    .counter-button:disabled {
        cursor: not-allowed;
    }

    .error {
        color: var(--vp-c-danger);
        background-color: var(--vp-c-danger-soft);
        padding: 0.75rem 1rem;
        border-radius: 0.5rem;
        text-align: center;
    }
</style>