document.getElementById('shorten-form').addEventListener('submit', async function (e) {
    e.preventDefault();

    const urlInput = document.getElementById('url');
    const url = urlInput.value;
    const resultDiv = document.getElementById('result');

    try {
        const response = await fetch('/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ url }),
        });

        if (!response.ok) {
            throw new Error('Failed to generate short URL');
        }

        const data = await response.json();
        resultDiv.innerHTML = `<p>Short URL: <a href="${data.short_url}" target="_blank">${data.short_url}</a></p>`;
    } catch (error) {
        resultDiv.innerHTML = `<p style="color: red;">Error: ${error.message}</p>`;
    }
});
