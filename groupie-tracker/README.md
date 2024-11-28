# Groupie-Tracker-Filters

A dynamic web application for music enthusiasts to discover and track their favorite bands and artists. This project provides an interactive interface to filter and search through artist information, concert dates, and locations.

## Features

### 🎵 Artist Information
- Comprehensive artist profiles including:
  - Band/Artist name
  - Formation year
  - First album release date
  - Band members
  - Profile images
  - Discography highlights

### 🗺️ Interactive Features
- **Advanced Filtering System:**
  - Filter by creation date (1950-2024)
  - Filter by first album release date
  - Filter by number of band members (1-8)
  - Filter by location/region
  - Real-time filter updates
  
- **Search Functionality:**
  - Dynamic search suggestions
  - Auto-complete feature
  - Search by artist name, location, or member

- **Interactive Map Integration:**
  - Powered by Leaflet.js
  - Visual representation of concert locations
  - Clickable markers with event details
  - Zoom and pan capabilities

### 📱 User Interface
- Responsive design for all devices
- Intuitive navigation
- Card-based artist display
- Filter and search icons for easy access
- Clean and modern aesthetic

## Technical Details

### API Structure
The application integrates with four main API endpoints:

1. **Artists API**
   - Basic artist information
   - Band member details
   - Formation dates
   - Album information

2. **Locations API**
   - Concert venues
   - Geographic coordinates
   - Region/country information

3. **Dates API**
   - Concert schedules
   - Tour dates
   - Event timing

4. **Relation API**
   - Data relationship mapping
   - Cross-referenced information
   - Unified data structure

### Technology Stack
- **Backend:** Go
- **Frontend:** HTML, CSS, JavaScript
- **Map Integration:** Leaflet.js
- **Data Format:** JSON

## Installation and Setup

### Prerequisites
- Go (latest stable version)
- Git
- Modern web browser
- Internet connection for map functionality

### Installation Steps

1. Clone the repository:
```bash
git clone https://learn.zone01kisumu.ke/git/moonyango/groupie-tracker-filters
```

2. Navigate to project directory:
```bash
cd groupie-tracker-filters
```

3. Start the server:
```bash
go run .
```

4. Access the application:
   - Open your web browser
   - Navigate to `http://localhost:8000`
   - Start exploring!

## Usage Guide

### Basic Navigation
1. Browse through artist cards on the main page
2. Click any artist card to view detailed information

### Using Filters
1. Click the filter icon in the navigation bar
2. Select your preferred criteria:
   - Adjust date ranges using sliders
   - Select number of members using checkboxes
   - Enter location in the search field
3. Click "Apply Filters" to update results
4. Use "Reset Filters" to clear all selections

### Search Function
1. Click the search icon or use the search bar
2. Start typing to see suggestions
3. Click on suggestions to view matching results

## Development and Contributing

### Contributing Guidelines
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to your fork
5. Submit a pull request

### Code Style
- Follow Go standard formatting
- Maintain consistent HTML/CSS structure
- Document new features thoroughly

## Project Team

<table>
<tr>
<td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
<a href=https://learn.zone01kisumu.ke/git/aaochieng>
<img src=https://learn.zone01kisumu.ke/git/avatars/8a1b24358854eb12998a07c269542193?size=870 width="100;" style="border-radius:50%" alt=Aaron/>
<br />
<sub style="font-size:14px"><b>Aaron Ochieng</b></sub>
</a>
</td>
<td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
<a href=https://learn.zone01kisumu.ke/git/moonyango>
<img src=https://learn.zone01kisumu.ke/git/avatars/8f9cf111e69139c3033e3b9f679e91ce?size=870 width="100;" style="border-radius:50%" alt=Moses/>
<br />
<sub style="font-size:14px"><b>Moses Onyango</b></sub>
</a>
</td>
<td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
<a href=https://learn.zone01kisumu.ke/git/bshisia>
<img src=https://learn.zone01kisumu.ke/git/avatars/e6b283b461701c3a0bde65d94393f768?size=870 width="100;" style="border-radius:50%" alt=Brian/>
<br />
<sub style="font-size:14px"><b>Brian Shisia</b></sub>
</a>
</td>
</tr>
</table>

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
- Leaflet.js for map integration
- Zone01 Kisumu for project support
- All contributors and testers

---
For additional information or support, please open an issue in the repository.

Happy exploring! 🎸✨
