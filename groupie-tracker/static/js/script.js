let currentPage = 1;
let itemsPerPage = 12;
let filteredCards = [];

function initializePagination() {
  const artistCards = document.querySelectorAll(".card-grid");
  filteredCards = Array.from(artistCards);

  currentPage = 1; // Set the initial page to 1
  updatePagination();
}

function updatePagination() {
  const totalPages = Math.ceil(filteredCards.length / itemsPerPage);
  document.getElementById("currentPage").textContent = currentPage;
  document.getElementById("totalPages").textContent = totalPages;

  document.getElementById("firstPage").disabled = currentPage === 1;
  document.getElementById("prevPage").disabled = currentPage === 1;
  document.getElementById("nextPage").disabled =
    currentPage === totalPages;
  document.getElementById("lastPage").disabled =
    currentPage === totalPages;

  // Show only the cards for the current page
  filteredCards.forEach((card, index) => {
    const shouldShow =
      index >= (currentPage - 1) * itemsPerPage &&
      index < currentPage * itemsPerPage;
    card.style.display = shouldShow ? "" : "none";
  });
}

function goToPage(action) {
  const totalPages = Math.ceil(filteredCards.length / itemsPerPage);

  switch (action) {
    case "first":
      currentPage = 1;
      break;
    case "prev":
      currentPage = Math.max(1, currentPage - 1);
      break;
    case "next":
      currentPage = Math.min(totalPages, currentPage + 1);
      break;
    case "last":
      currentPage = totalPages;
      break;
  }

  updatePagination();
}

function updateItemsPerPage() {
  itemsPerPage = parseInt(document.getElementById("itemsPerPage").value);
  currentPage = 1;
  updatePagination();
}

function filterArtists() {
  const searchInput = document.getElementById("searchInput");
  const filter = searchInput.value.toLowerCase();
  const artistCards = document.querySelectorAll(".card-grid");

  filteredCards = Array.from(artistCards).filter((card) => {
    const artistName = card.getAttribute("data-name").toLowerCase();
    return artistName.includes(filter);
  });

  currentPage = 1;
  updatePagination();
}
function initializeKeyboardNavigation() {
  const gridItems = document.querySelectorAll(".card-grid");
  let currentFocusIndex = -1;

  // Function to focus on a grid item
  function focusGridItem(index) {
    // Remove focus from all items
    gridItems.forEach((item) => item.setAttribute("tabindex", "-1"));

    // Ensure index is within bounds
    if (index >= 0 && index < gridItems.length) {
      const targetItem = gridItems[index];

      // Only focus visible items
      if (targetItem.style.display !== "none") {
        targetItem.setAttribute("tabindex", "0");
        targetItem.focus();
        currentFocusIndex = index;

        // If the focused item is not in the current page, go to its page
        const itemPage = Math.floor(index / itemsPerPage) + 1;
        if (itemPage !== currentPage) {
          currentPage = itemPage;
          updatePagination();
        }
      }
    }
  }

  // Add tabindex to all grid items
  gridItems.forEach((item, index) => {
    item.setAttribute("tabindex", "-1");

    // Add click listener to update currentFocusIndex
    item.addEventListener("click", () => {
      currentFocusIndex = index;
    });
  });

  // Add keyboard event listener to the document
  document.addEventListener("keydown", (e) => {
    // Only handle keyboard navigation when not in an input field
    if (e.target.tagName === "INPUT") return;

    const visibleItems = Array.from(gridItems).filter(
      (item) => item.style.display !== "none"
    );
    const itemsPerRow = Math.floor(
      document.querySelector(".artists").offsetWidth /
        gridItems[0].offsetWidth
    );

    switch (e.key) {
      case "ArrowRight":
        e.preventDefault();
        if (currentFocusIndex === -1) {
          focusGridItem(0);
        } else {
          focusGridItem(currentFocusIndex + 1);
        }
        break;
      case "ArrowLeft":
        e.preventDefault();
        if (currentFocusIndex > 0) {
          focusGridItem(currentFocusIndex - 1);
        }
        break;
      case "ArrowDown":
        e.preventDefault();
        if (currentFocusIndex === -1) {
          focusGridItem(0);
        } else {
          focusGridItem(currentFocusIndex + itemsPerRow);
        }
        break;
      case "ArrowUp":
        e.preventDefault();
        if (currentFocusIndex >= itemsPerRow) {
          focusGridItem(currentFocusIndex - itemsPerRow);
        }
        break;
      case "Enter":
        if (currentFocusIndex !== -1) {
          const link = gridItems[currentFocusIndex].querySelector("a");
          if (link) link.click();
        }
        break;
    }
  });
}

document.addEventListener("DOMContentLoaded", () => {
  const loadingScreen = document.getElementById("loading-screen");
  const container = document.querySelector(".container");
  // Simulate loading time (you can adjust this based on your actual data fetching)
  setTimeout(() => {
    loadingScreen.style.opacity = "0";
    container.style.opacity = "1";
    setTimeout(() => {
      loadingScreen.style.display = "none";
    }, 500); // This matches the transition time
  }, 2000); // Adjust this time as needed
  initializePagination();
  initializeKeyboardNavigation();
});

function filterArtists() {
  const searchInput = document.getElementById("searchInput");
  const filter = searchInput.value.toLowerCase();
  const artistCards = document.querySelectorAll(".card-grid");

  artistCards.forEach((card) => {
    const artistName = card.getAttribute("data-name").toLowerCase();
    if (artistName.includes(filter)) {
      card.style.display = "";
    } else {
      card.style.display = "none";
    }
  });
}

document.addEventListener("DOMContentLoaded", () => {
  const artistCards = document.querySelectorAll(".card-grid");
  artistCards.forEach((card) => {
    card.style.display = "";
  });
});