document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the default form submission

    // Get the values from the form
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const userType = document.getElementById('userType').value;

    // Here you would typically validate the username and password with a backend service
    // For this example, we will just redirect based on userType

    if (userType === 'admin') {
        window.location.href = 'admin_dashboard.html'; // Redirect to admin dashboard
    } else if (userType === 'waste_owner') {
        window.location.href = 'waste_owner_dashboard.html'; // Redirect to waste owner dashboard
    } else if (userType === 'recycler') {
        window.location.href = 'recycler_dashboard.html'; // Redirect to recycler dashboard
    } else {
        alert('Please select a valid user type.');
    }
});
