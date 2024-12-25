export const ssr = false;
export const prerender = false;


export async function load() {
    try {
        const response = await fetch('/api/count');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return {
            count: data.count
        };
    } catch (error) {
        console.error('Error fetching count:', error);
        return {
            count: 0
        };
    }
}