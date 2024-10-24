// WebSocket bağlantısını kur
const socket = new WebSocket('ws://localhost:3000/ws'); // WebSocket sunucunuzun URL'sini buraya ekleyin

// WebSocket bağlantısı açıldığında
socket.onopen = function(event) {
    console.log('WebSocket is open now.');
};

// WebSocket'ten veri alındığında
socket.onmessage = function(event) {
    const data = JSON.parse(event.data);
    console.log('Data received:', data);

    // WebSocket'ten alınan verileri güncelle
    updateTicketList(data);
};

// WebSocket bağlantısı kapandığında
socket.onclose = function(event) {
    console.log('WebSocket is closed now.');
};

// WebSocket bağlantısında bir hata oluştuğunda
socket.onerror = function(error) {
    console.error('WebSocket Error:', error);
};

// Listeyi güncelle
function updateTicketList(data) {
    const container = document.getElementById('ticket-cards');
    container.innerHTML = '';

    // Türler için varsayılan boş dizi
    const wheelchairTickets = data.wheelchair || [];
    const childTickets = data.child || [];
    const personTickets = data.person || [];

    // Tüm biletleri birleştir ve türlerine göre sıralayın
    const tickets = [
        ...wheelchairTickets.map(ticket => ({ ...ticket, type: 'wheelchair' })),
        ...childTickets.map(ticket => ({ ...ticket, type: 'child' })),
        ...personTickets.map(ticket => ({ ...ticket, type: 'person' }))
    ];

    // Türlerine göre sıralayın
    const sortedTickets = tickets.slice().sort((a, b) => {
        if (a.type === 'wheelchair' && b.type !== 'wheelchair') return -1;
        if (a.type === 'child' && b.type === 'person') return -1;
        if (a.type === 'person' && b.type === 'wheelchair') return 1;
        if (a.type === 'wheelchair' && b.type === 'child') return -1;
        if (a.type === 'child' && b.type === 'wheelchair') return 1;
        return 0;
    });

    if (sortedTickets.length === 0) {
        container.innerHTML = '<p>Nobatda adam yok.</p>';
    } else {
        sortedTickets.forEach((ticket, index) => {
            const card = document.createElement('div');
            card.className = 'card';
            card.style.backgroundColor = getCardColor(index);
            card.innerHTML = `                
                <h1>${ticket.ticket}</h1>
                <hr>
                <p class="title">${ticket.type}</p>
            `;
            container.appendChild(card);
        });
    }
}

// Kart rengini indeksine göre belirle
function getCardColor(index) {
    if (index === 0) return 'green' ;
    if (index === 1) return 'orange';
    return 'red';
}
