Toy project to create an image downloader for the Netrunner cardgame on OCTGN

Downloads uncensored card images to be used by the application.

Features:
  Tries multiple sites before fail
  No updates required to dl new set releases

Current Goals:
  Only download image if theres a higher rez one available
  More failback sites
  Multithreaded
  Reuse connection for multiple downloads
  Rewrite into Go (for learning/multiprocessing)
  