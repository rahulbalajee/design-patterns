package main

import "github.com/rahulbalajee/design-patterns/streamer"

const numWorkers = 4

// videoQueue is package-level but private
var videoQueue chan streamer.VideoProcessingJob

// initStreamer initializes the video processing queue
func initStreamer() {
	videoQueue = make(chan streamer.VideoProcessingJob, numWorkers)
}

// closeStreamer properly closes the video queue
func closeStreamer() {
	if videoQueue != nil {
		close(videoQueue)
	}
}
