# Bookings and Reservations

# Section 17: Setting up secure back end administration

# 143. Important: A note on the admin.layout.tmpl file

Important: A note on the admin.layout.tmpl file
When I recorded this lecture, the version of bootstrap used by the admin template was 5.0.x (I'm not sure about the last digit). Since then, the admin template has been updated to a newer version of bootstrap which requires one change for the left-hand navigation. Anywhere you see a line with something like this:

<a class="nav-link" data-toggle="collapse"...

Please change it to this:

<a class="nav-link" data-bs-toggle="collapse" ...

Note the change from data-toggle to data-bs-toggle

If you fail to make this change, your collapsible menu item in the left-hand navigation will not expand as it should.
