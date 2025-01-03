document.addEventListener("DOMContentLoaded", () => {
    const popup = document.getElementById("reservation-popup");
    const closePopup = document.getElementById("close-popup");
    const reservationForm = document.getElementById("reservation-form");
    let currentHouseId = null;

    // Show popup
    document.querySelectorAll(".tile button").forEach(button => {
        button.addEventListener("click", event => {
            const houseId = event.target.dataset.houseId;
            currentHouseId = houseId; // Store current house ID
            popup.style.display = "block";
        });
    });

    // Close popup
    closePopup.addEventListener("click", () => {
        popup.style.display = "none";
    });

    // Submit form
    reservationForm.addEventListener("submit", async (event) => {
        event.preventDefault();

        const name = document.getElementById("name").value;
        const email = document.getElementById("email").value;
        const startDate = document.getElementById("start-date").value;
        const endDate = document.getElementById("end-date").value;

        const response = await fetch("/reserve", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                house_id: currentHouseId,
                name,
                email,
                start_date: startDate,
                end_date: endDate,
            }),
        });

        if (response.ok) {
            alert("Reservation successful!");
            popup.style.display = "none";
        } else {
            alert("Error making reservation. Please try again.");
        }
    });
});
