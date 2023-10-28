desc "Builds relinkvc for release"

Envs = [
  { goos: "darwin", arch: "386" },
  { goos: "darwin", arch: "amd64" },
  { goos: "linux", arch: "arm" },
  { goos: "linux", arch: "arm64" },
  { goos: "linux", arch: "386" },
  { goos: "linux", arch: "amd64" },
  { goos: "windows", arch: "386" },
  { goos: "windows", arch: "amd64" }
].freeze

Version = "0.5.0".freeze

task :build do
  `rm -rf dist/#{Version}`
  Envs.each do |env|
    ENV["GOOS"] = env[:goos]
    ENV["GOARCH"] = env[:arch]
    puts "Building #{env[:goos]} #{env[:arch]}"
    `GOOS=#{env[:goos]} GOARCH=#{env[:arch]} CGO_ENABLED=0 go build -v -o dist/#{Version}/relinkvc`
    if env[:goos] == "windows"
      puts "Creating windows executable"
      `mv dist/#{Version}/relinkvc dist/#{Version}/relinkvc.exe`
      `cd dist/#{Version} && zip relinkvc_win.zip relinkvc.exe`
      puts "Removing windows executable"
      `rm -rf dist/#{Version}/relinkvc.exe`
    else
      puts "Tarring #{env[:goos]} #{env[:arch]}"
      `cd dist/#{Version} && tar -czvf relinkvc_#{env[:goos]}_#{env[:arch]}.tar.gz relinkvc`
      puts "Removing dist/#{Version}/relinkvc"
      `rm -rf dist/#{Version}/relinkvc`
    end
  end
end

task default: :build
