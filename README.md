Afrostream Media Server (AMS)
===
Afrostream Media Server is a streaming software implemented in [Go](http://golang.org).

### Synopsis
Afrostream Media Server (AMS) is an OpenSource software for streaming MP4 audio/video files to various formats (like **DASH**, **HLS** and **Smooth Streaming**). Currently, the 0.1-alpha version only supports DASH, The implementation of HLS and Smooth Streaming is underway. The goal of this project is to provide an **Unified Streaming (c)** like OpenSource software. Feel free to contact and/or join us to participate to this great project. AMS is considered as experimental project.

### How to build
AMS is developped in GoLang so you can build the software for any Operation System you want.

First, clone the github project to a local directory:

	git clone https://github.com/Afrostream/afrostream-media-server.git

Enter to the directory freshly created by git:

	cd afrostream-media-server

For the next step, you will need [Go](http://golang.org) installed, please refer to your OS for installing Go.

Build amspackager (the packager) and ams (the media server), replace <fullpath> with the base directory where you ran git clone command:

	export GOPATH=<fullpath>/afrostream-media-server
	go build amspackager.go
	go build ams.go

Install binaries in a bin directory (eg: /usr/local/bin or /usr/bin):

	cp amspackager /usr/local/bin/
	cp ams /usr/local/bin/

Now you've two binaries amspackager and ams installed on your OS.

### How to run
The first thing to do is creating mp4 files needed by amspackager and ams to create DASH fragments. We use FFMpeg software to generate these mp4 files.

First, We create 4 video profiles and 1 audio profile with a video.mp4 files containing 1 video stream and 1 audio stream:

	ffmpeg -i video.mp4 -y -vf yadif=0:-1:0,scale="426:trunc(ow/a/2)*2",setsar=1:1 -c:v libx264 -preset:v fast -profile:v baseline -level 3.0 -coder:v 0 -b:v 400k -minrate 400k -maxrate 400k -bufsize 400k -g 50 -keyint_min 50 -sc_threshold 0 -pix_fmt yuv420p -map 0:v -an -map_chapters -1 -threads 0 video_h264-426x240-400.mp4
	ffmpeg -i video.mp4 -y -vf yadif=0:-1:0,scale="640:trunc(ow/a/2)*2",setsar=1:1 -c:v libx264 -preset:v fast -profile:v baseline -level 3.0 -coder:v 0 -b:v 800k -minrate 800k -maxrate 800k -bufsize 800k -g 50 -keyint_min 50 -sc_threshold 0 -pix_fmt yuv420p -map 0:v -an -map_chapters -1 -threads 0 video_h264-640x360-800.mp4
	ffmpeg -i video.mp4 -y -vf yadif=0:-1:0,scale="854:trunc(ow/a/2)*2",setsar=1:1 -c:v libx264 -preset:v fast -profile:v baseline -level 3.1 -coder:v 1 -b:v 1600k -minrate 1600k -maxrate 1600k -bufsize 1600k -g 50 -keyint_min 50 -sc_threshold 0 -pix_fmt yuv420p -map 0:v -an -map_chapters -1 -threads 0 video_h264-854x480-1600.mp4
	ffmpeg -i video.mp4 -y -vf yadif=0:-1:0,scale="1280:trunc(ow/a/2)*2",setsar=1:1 -c:v libx264 -preset:v fast -profile:v main -level 3.1 -coder:v 1 -b:v 3000k -minrate 3000k -maxrate 3000k -bufsize 3000k -g 50 -keyint_min 50 -sc_threshold 0 -pix_fmt yuv420p -map 0:v -an -map_chapters -1 -threads 0 video_h264-1280x720-3000.mp4
	ffmpeg -i video.mp4 -y -vn -acodec libfaac -ac 2 -ab 128k -ar 48000 -map 0:a:0 -map_chapters -1 -threads 0 video_aac-128.mp4

Output profiles will be:

	Video H264 @ 426x240  400kbits/s  Baseline profile (3.0)
	Video H264 @ 640x360  800kbits/s  Baseline profile (3.0)
	Video H264 @ 854x480  1600kbits/s Baseline profile (3.1)
	Video H264 @ 1280x720 3000kbits/s Main profile (3.1)
	Audio AAC  @ 48000khz 128kbits/s

Move all mp4 files to a directory that you'll use for the HTTP media server document root, cd to this directory and run amspackager to prepare the content for AMS:

	/usr/local/bin/amspackager -o video.json -d 8 -i video_h264-426x240-400.mp4 -i video_h264-640x360-800.mp4 -i video_h264-854x480-1600.mp4 -i video_h264-1280x720-3000.mp4 -i video_aac-128.mp4

Output will be:

	AMSPackager -- spebsd@gmail.com / Afrostream
	
	-- Parsing file='video_h264-426x240-400.mp4' language='eng'
	-- Parsing file='video_h264-640x360-800.mp4' language='eng'
	-- Parsing file='video_h264-854x480-1600.mp4' language='eng'
	-- Parsing file='video_h264-1280x720-3000.mp4' language='eng'
	-- Parsing file='video_aac-128.mp4' language='eng'
	
	-- Creating package file 'video.json'
	
	All files has been packaged successfully

If you have vtt subtitles files, you can add them with -i video.en.vtt -l eng -i video.fr.vtt -l fra ...
Your video is prepared for AMS, so let's run Afrostream Media Server as root and listening on HTTP port 80 (you can package any video files on the fly without restarting AMS):

	# /usr/local/bin/ams -d <document_root_path> -p 80

Now, you can request URL http://<ip_of_your_server>/video.json/.mpd with a dash player like [DASHJS](http://dashif.org/reference/players/javascript/v1.5.1/samples/dash-if-reference-player/index.html). That's all.

## TODO
<table>
<tr>
<th>Functionnality</th>
<th>Implemented</th>
</tr>
<tr>
<th>DASH on-the-fly</th>
<th>Yes</th>
</tr>
<tr>
<th>HLS on-the-fly</th>
<th>No</th>
</tr>
<tr>
<th>Smooth Streaming on-the-fly</th>
<th>No</th>
</tr>
<tr>
<th>DRM</th>
<th>No</th>
</tr>
</table>

## Bugs
There is probably some bugs in this implementation, all mp4 parsing has been developped in pure GoLang, no external libraries has been used. If you see any bugs please report to tech@afrostream.tv or spebsd@gmail.com.
