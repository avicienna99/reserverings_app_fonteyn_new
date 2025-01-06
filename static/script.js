document.addEventListener("DOMContentLoaded", () => {
    const popup = document.getElementById("popup");
    const overlay = document.getElementById("overlay");
    const closePopupButton = document.getElementById("close-popup"); // Renamed to avoid conflict
    const reservationForm = document.getElementById("reservation-form");

    // Log DOM elements to confirm they're loaded
    console.log("Popup element:", popup);
    console.log("Overlay element:", overlay);
    console.log("Reservation Form:", reservationForm);

    // Load house tiles
    const houses = document.querySelectorAll(".house");
    console.log("Found house elements:", houses);

    houses.forEach(house => {
        console.log("Attaching event listener to house:", house);
        house.addEventListener("click", () => {
            console.log(`House clicked: ID = ${house.dataset.id}`);
            openPopup(house.dataset.id);
        });
    });

    function openPopup(houseId) {
        console.log(`Opening popup for house ID: ${houseId}`);
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
        console.log("Closing popup");
        popup.classList.remove("visible");
        overlay.classList.remove("visible");
    }

    closePopupButton.addEventListener("click", closePopup); // Updated usage
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

        console.log("Submitting reservation:", reservation);

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
