require 'fileutils'

def echo_system(cmd)
  puts("+ #{cmd}")
  system(cmd)
end

def backup()
  # Define target directory path
  # TODO: Hard code
  target_dir_path = "./mydir"
  # TODO: Hard code
  dst_dir_path = "./mydst"
  # Get base name of the target directory
  target_basename = File.basename(target_dir_path)
  # Define snar name
  target_snar_name = "#{target_basename}.snar"

  # Create destination path if it doesn't exist
  FileUtils.mkdir_p(dst_dir_path)

  # Get current time
  time = Time.new
  # Create string representation of current time
  time_str = time.strftime("%Y%m%d_%H%M_%S_") + time.zone
  # Define tar file name
  tar_file_name = "#{target_basename}_#{time_str}.tar"


  target_snar_path = File.join(dst_dir_path, target_snar_name)
  tar_file_path = File.join(dst_dir_path, tar_file_name)

  # Incremental backup
  echo_system("gtar -g #{target_snar_path} -cf #{tar_file_path} #{target_dir_path}")
end


def restore()
  # TODO: Hard code
  dst_dir_path = "./mydst"
  Dir.chdir(dst_dir_path) {
    # Get snar file name
    snar_file_name = Dir.glob("**.snar").first
    # Get tar names
    tar_file_names = Dir.glob("**.tar").sort
    # Restore
    tar_file_names.each{|tar_file_name|
      echo_system("gtar -g #{snar_file_name} -xf #{tar_file_name}")
    }
  }
end

# backup()
restore()