function handleClick(type) {
    showModal();

    // URL'yi query parametresi ile belirleyin
    let url = `http://localhost:3000?type=${encodeURIComponent(type)}`;

    // GET isteği gönder
    fetch(url)
        .then(response => response.json())
        .then(data => {
            console.log(data); 
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



