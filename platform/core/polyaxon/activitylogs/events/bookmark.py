import activitylogs

from events.registry import bookmark

activitylogs.subscribe(bookmark.BookmarkBuildJobsViewedEvent)
activitylogs.subscribe(bookmark.BookmarkJobsViewedEvent)
activitylogs.subscribe(bookmark.BookmarkExperimentsViewedEvent)
activitylogs.subscribe(bookmark.BookmarkExperimentGroupsViewedEvent)
activitylogs.subscribe(bookmark.BookmarkProjectsViewedEvent)
