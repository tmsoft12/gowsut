// WebSocket bağlantısını kur
const socket = new WebSocket('ws://192.168.100.224:8000/ws'); // WebSocket sunucunuzun URL'sini buraya ekleyin

// WebSocket bağlantısı açıldığında
socket.onopen = function (event) {
    console.log('WebSocket açyldy.');
};

// WebSocket'ten veri alındığında
socket.onmessage = function (event) {
    const data = JSON.parse(event.data);
    console.log('Alnan maglumat:', data);

    // WebSocket'ten alınan verileri güncelle
    updateTicketList(data);
};

// WebSocket bağlantısı kapandığında
socket.onclose = function (event) {
    console.log('WebSocket ýapylýar.');
};

// WebSocket bağlantısında bir hata oluştuğunda
socket.onerror = function (error) {
    console.error('WebSocket Hatası:', error);
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
        container.innerHTML = '<p>Nobatda adam ýok.</p>';
    } else {
        sortedTickets.forEach((ticket, index) => {
            const card = document.createElement('div');
            card.className = 'card';
            const cardColor = getCardColor(index);
            card.style.backgroundColor = cardColor; // Arka plan rengi (yumuşak ton)
            card.style.border = `2px solid ${getBorderColor(cardColor)}`; // Kenar rengi (koyu ton)
            card.innerHTML = `
                <div class="card-content">
                    <h1 class="ticket-title">${ticket.ticket}</h1>
                    <hr>
                    <p class="title">${ticket.type}</p>
                </div>
            `;
            container.appendChild(card);
        });
    }
}

// Kart rengini indeksine göre belirle (yumuşak tonlar)
function getCardColor(index) {
    if (index === 0) return '#e8f5e9';  // Yumuşak yeşil
    if (index === 1) return '#fff3e0';  // Yumuşak turuncu
    return '#ffebee';                  // Yumuşak kırmızı
}

// Kenar rengini kart renginden ayarlayın (koyu tonlar)
function getBorderColor(cardColor) {
    if (cardColor === '#e8f5e9') return '#388e3c';   // Koyu yeşil
    if (cardColor === '#fff3e0') return '#ff9800';   // Koyu turuncu
    return '#d32f2f';                               // Koyu kırmızı
}
