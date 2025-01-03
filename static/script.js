document.addEventListener("DOMContentLoaded", () => {
    const popup = document.getElementById("popup");
    const overlay = document.getElementById("overlay");
    const closePopup = document.getElementById("close-popup");
    const reservationForm = document.getElementById("reservation-form");

    // Load house tiles
    document.querySelectorAll(".house").forEach(house => {
        house.addEventListener("click", () => {
            openPopup(house.dataset.id);
        });
    });

    function openPopup(houseId) {
        popup.classList.add("visible");
        overlay.classList.add("visible");
        reservationForm.dataset.houseId = houseId;

        // Clear form inputs
        document.getElementById("name").value = '';
        document.getElementById("email").value = '';
        document.getElementById("start-date").value = '';
        document.getElementById("end-date").value = '';
    }

    function closePopup() {
        popup.classList.remove("visible");
        overlay.classList.remove("visible");
    }

    closePopup.addEventListener("click", closePopup);
    overlay.addEventListener("click", closePopup);

    reservationForm.addEventListener("submit", async (e) => {
        e.preventDefault();
        const houseId = reservationForm.dataset.houseId;
        const name = document.getElementById("name").value;
        const email = document.getElementById("email").value;
        const startDate = document.getElementById("start-date").value;
        const endDate = document.getElementById("end-date").value;

        const reservation = {
            house_id: parseInt(houseId),
            name,
            email,
            start_date: startDate,
            end_date: endDate,
        };

        try {
            const response = await fetch('/reserve', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(reservation),
            });

            if (response.ok) {
                alert('Reservation successful!');
            } else {
                alert('Failed to submit reservation.');
            }
        } catch (error) {
            console.error('Error submitting reservation:', error);
            alert('An error occurred. Please try again.');
        }

        closePopup();
    });
});
