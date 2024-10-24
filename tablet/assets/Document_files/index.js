
function handleClick(type) {
    showModal();

    // URL'yi belirleyin, örneğin:
    let url = `https://example.com/api/${type}`;

    // GET isteği gönder
    fetch(url)
        .then(response => response.json())
        .then(data => {
            console.log(data); // Gelen veriyi işleyin
            // Modal 3 saniye sonra kapanacak
            setTimeout(function() {
                closeModal();
            }, 3000);
        })
        .catch(error => {
            console.error('Error:', error);
            setTimeout(function() {
                closeModal();
            }, 3000);
        });
}

function showModal() {
    document.getElementById('myModal').style.display = 'flex';
}

function closeModal() {
    document.getElementById('myModal').style.display = 'none';
}

// Modal dışına tıklanıldığında kapatma
window.onclick = function(event) {
    if (event.target === document.getElementById('myModal')) {
        closeModal();
    }
}
