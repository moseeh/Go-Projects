const cardsPerPage = 18; // Number of artist cards per page
let currentPage = 1; // Current page
const artistCards = document.querySelectorAll(".artist-card");
const totalPages = Math.ceil(artistCards.length / cardsPerPage);
// Display the current page of artist cards
const showPage = (page) => {
  const start = (page - 1) * cardsPerPage;
  const end = page * cardsPerPage;
  artistCards.forEach((card, index) => {
    card.style.display = index >= start && index < end ? "block" : "none";
  });
  document.getElementById(
    "page-info"
  ).textContent = `Page ${page} of ${totalPages}`;
  document.getElementById("prev-page").disabled = page === 1;
  document.getElementById("next-page").disabled = page === totalPages;
};
// Change page function (prev or next)
const changePage = (direction) => {
  currentPage += direction;
  showPage(currentPage);
};
// Initially display the first page
showPage(currentPage);
const filterArtists = () => {
  const searchQuery = document.getElementById("search-bar").value.toLowerCase();
  const container = document.querySelector(".suggestions");
  container.innerHTML = "";
  if (searchQuery === "") {
    container.style.display = "none";
    showPage(currentPage);
    return;
  }
  container.style.display = "block";
  let suggestionsAdded = false;
  artistCards.forEach((card) => {
    const artistName = card.querySelector("h2").textContent;
    const artistId = card.getAttribute("data-id");
    const artistMembers = card
      .getAttribute("data-members")
      .split(",")
      .map((member) => member.trim())
      .filter((member) => member !== ""); // Filter out empty strings
    const artistLocations = card
      .getAttribute("data-locations")
      .split(",")
      .map((location) => location.trim())
      .filter((location) => location !== ""); // Filter out empty strings
    const artistFirstAlbum = card.getAttribute("data-firstalbum");
    const artistCreationDate = card.getAttribute("data-creationdate");
    // Check artist name
    if (artistName.toLowerCase().includes(searchQuery)) {
      createSuggestion(container, `${artistName} - artist/band`, artistId);
      suggestionsAdded = true;
    }
    // Check creation date
    if (artistCreationDate.toLowerCase().includes(searchQuery)) {
      createSuggestion(
        container,
        `${artistCreationDate} - creation date of ${artistName}`,
        artistId
      );
      suggestionsAdded = true;
    }
    // Check first album
    if (artistFirstAlbum.toLowerCase().includes(searchQuery)) {
      createSuggestion(
        container,
        `${artistFirstAlbum} - first album of ${artistName}`,
        artistId
      );
      suggestionsAdded = true;
    }
    // Check members
    artistMembers.forEach((member) => {
      if (member.toLowerCase().includes(searchQuery)) {
        createSuggestion(
          container,
          `${member} - member of ${artistName}`,
          artistId
        );
        suggestionsAdded = true;
      }
    });
    // Check locations
    artistLocations.forEach((location) => {
      if (location.toLowerCase().includes(searchQuery)) {
        createSuggestion(
          container,
          `${location} - location of ${artistName}`,
          artistId
        );
        suggestionsAdded = true;
      }
    });
  });
  if (!suggestionsAdded) {
    container.innerHTML = "<p>No matches found</p>";
  }
};
const createSuggestion = (container, textContent, artistId) => {
  const button = document.createElement("button");
  button.textContent = textContent;
  button.className = "suggestion-button";
  button.addEventListener("click", () => {
    window.location.href = `/details/${artistId}`;
  });
  container.appendChild(button);
};
// Close suggestions when clicking outside
document.addEventListener("click", function (event) {
  const container = document.querySelector(".suggestions");
  const searchBar = document.getElementById("search-bar");
  if (!container.contains(event.target) && event.target !== searchBar) {
    container.style.display = "none";
  }
});

// Toggle filter menu visibility
function toggleFilters() {
  document.getElementById("filterMenu").classList.toggle("active");
}

// Close filters when clicking outside
document.addEventListener("click", (event) => {
  const filterMenu = document.getElementById("filterMenu");
  const filterButton = event.target.closest(".filter-button");
  const filterMenuElement = event.target.closest(".filter-menu");

  if (
    !filterButton &&
    !filterMenuElement &&
    filterMenu.classList.contains("active")
  ) {
    filterMenu.classList.remove("active");
  }
});

// Add this to your initialization
document.addEventListener("DOMContentLoaded", () => {
  initializeArtists();
  populateYearFilter();
});


function setupRangeInputs(startInput, endInput, startDisplay, endDisplay) {
  function updateRanges() {
      let startVal = parseInt(startInput.value);
      let endVal = parseInt(endInput.value);

      // Ensure end value is always greater than start value
      if (startVal >= endVal) {
          if (this === startInput) {
              endInput.value = Math.min(parseInt(startVal) + 1, endInput.max);
              endVal = parseInt(endInput.value);
          } else {
              startInput.value = Math.max(parseInt(endVal) - 1, startInput.min);
              startVal = parseInt(startInput.value);
          }
      }

      // Update displays
      startDisplay.textContent = startVal;
      endDisplay.textContent = endVal;
  }

  startInput.addEventListener('input', updateRanges);
  endInput.addEventListener('input', updateRanges);
}

// Setup Creation Date Range
setupRangeInputs(
  document.querySelector('input[name="creationDateStart"]'),
  document.querySelector('input[name="creationDateEnd"]'),
  document.getElementById('creationDateStartValue'),
  document.getElementById('creationDateEndValue')
);

// Setup First Album Range
setupRangeInputs(
  document.querySelector('input[name="firstAlbumStart"]'),
  document.querySelector('input[name="firstAlbumEnd"]'),
  document.getElementById('firstAlbumStartValue'),
  document.getElementById('firstAlbumEndValue')
);


// Function to update range display values
function updateRangeDisplays() {
  document.getElementById('creationDateStartValue').textContent = document.querySelector('input[name="creationDateStart"]').value;
  document.getElementById('creationDateEndValue').textContent = document.querySelector('input[name="creationDateEnd"]').value;
  document.getElementById('firstAlbumStartValue').textContent = document.querySelector('input[name="firstAlbumStart"]').value;
  document.getElementById('firstAlbumEndValue').textContent = document.querySelector('input[name="firstAlbumEnd"]').value;
}

// Function to reset all filters
function resetFilters(event) {
  // Prevent form submission
  event.preventDefault();
  
  // Get the form element
  const form = document.getElementById('filterMenu');
  
  // Reset range inputs to default values
  form.querySelector('input[name="creationDateStart"]').value = "1950";
  form.querySelector('input[name="creationDateEnd"]').value = "2024";
  form.querySelector('input[name="firstAlbumStart"]').value = "1950";
  form.querySelector('input[name="firstAlbumEnd"]').value = "2024";
  
  // Reset all checkboxes
  const checkboxes = form.querySelectorAll('input[type="checkbox"]');
  checkboxes.forEach(checkbox => {
    checkbox.checked = false;
  });
  
  // Reset location input
  form.querySelector('input[name="location"]').value = "";
  
  // Update the display values for ranges
  updateRangeDisplays();
}

// Add event listeners for range inputs to update displays
document.addEventListener('DOMContentLoaded', function() {
  const rangeInputs = document.querySelectorAll('input[type="range"]');
  rangeInputs.forEach(input => {
    input.addEventListener('input', updateRangeDisplays);
  });
  
  // Initial update of range displays
  updateRangeDisplays();
  
  // Add event listener to reset button
  const resetButton = document.querySelector('.reset-filters');
  resetButton.addEventListener('click', resetFilters);
});